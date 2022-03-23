"""empty message

Revision ID: db5d8e00ddd1
Revises: d5f1e6f26fd4
Create Date: 2022-03-22 18:24:08.991313

"""
import sqlalchemy as sa
from alembic import op

# revision identifiers, used by Alembic.
revision = 'db5d8e00ddd1'
down_revision = 'd5f1e6f26fd4'
branch_labels = None
depends_on = None


def upgrade():
    """
    Run upgrade
    """
    # ### commands auto generated by Alembic - please adjust! ###
    op.add_column(
        'cbsds', sa.Column(
            'preferred_bandwidth_mhz',
            sa.Integer(), server_default='0', nullable=False,
        ),
    )
    op.add_column(
        'cbsds', sa.Column(
            'preferred_frequencies_mhz',
            sa.JSON(), server_default='[]', nullable=False,
        ),
    )
    # ### end Alembic commands ###


def downgrade():
    """
    Run downgrade
    """
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_column('cbsds', 'preferred_frequencies_mhz')
    op.drop_column('cbsds', 'preferred_bandwidth_mhz')
    # ### end Alembic commands ###
